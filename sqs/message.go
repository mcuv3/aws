package sqs

import (
	"context"
	"time"

	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)




func (s *SQSServer) SendMessage(ctx context.Context,req * aws.SendMessageRequest)(*aws.SendMessageResponse,error){ 

	us ,err := s.auth.GetUserMetadata(ctx)


	if err != nil {
		return nil,s.Error(err,codes.Unauthenticated,"Unable to authenticate")
	}

	queue:= model.Queue{}

	res := s.db.Where("account_id = ? AND name = ?",us.AccountId,req.GetQueueName()).First(&queue)

	if res.Error !=nil {
		return nil,s.Error(res.Error,codes.NotFound,"Queue not found")
	}


	msg := model.Message{
		Message: req.GetMessage(),
		QueueID: queue.ID,
		UserMessageId: model.UserMessageId(uuid.NewString()),
		QueueName: queue.Name,
	}

	
	res = s.db.Create(&msg) 
 
	if res.Error != nil {	
		return nil,s.Error(res.Error,codes.InvalidArgument,"Unable to send the message.")
	}

	//	err = redis.Set(ctx, msg.Key(), m, 0).Err()

		// if res.Error != nil || err != nil {
		// 	return nil,s.Error() : "Could not save the message", Status: 400})
		// }
	return &aws.SendMessageResponse{
		Message: &aws.Message{
			Id: string(msg.UserMessageId),
			Message: msg.Message,
			QueueName: queue.Name,
		},
	},nil
}

func (s *SQSServer) ReceiveMessage(req *aws.ReceiveMessageRequest,stream aws.SQSService_ReceiveMessageServer)error{ 


	us , err := s.auth.Authorize(stream.Context(),"")
	
	if err != nil {	
		s.logger.Err(err).Str("error:",err.Error())
		return s.Error(err,codes.Unauthenticated,"Unable to authenticate")
	}
	
	queue := model.Queue{}

	res := s.db.Where("account_id = ? AND name = ?",us.AccountId,req.GetQueueName()).First(&queue)

	if res.Error !=nil {
		return s.Error(res.Error,codes.NotFound,"Queue not found")
	}



	// msgs := make(chan []model.Message)
	errChan := make(chan error)

		go func() {
			
			var sent,wait int 

			process:for {
				msgs := []*model.Message{}
 
				// keys, err := redis.Keys(ctx, queue.Pattern()).Result()

				 res := s.db.Where("queue_id = ? AND status = ?",queue.ID,model.Available).Find(&msgs) 

				 if res.Error !=nil {
					  errChan <- s.Error(res.Error,codes.NotFound,"Something went wrong getting the messages.")
				} 
				
				for _,m := range msgs {

					if sent == int(req.GetBatchLimit()) { break process }

					err := stream.SendMsg(&aws.ReceiveMessageResponse{Message: 
						&aws.Message{Id:  string(m.UserMessageId), Message: m.Message, QueueName: queue.Name}})
						
						if err !=nil {
							errChan <- err
					}
					m.Status = model.Processing
					m.DeleteOn = time.Now().Add(time.Duration(queue.Configuration.VisibilityTimeout) * time.Second)
					if err := s.db.Save(&m).Error; err !=nil {
						errChan <- err 
					}
					sent++
				}
					

				// if err == nil {
				// 	for _, k := range keys {
				// 		val, _ := redis.Get(ctx, k).Result()
				// 		redis.Del(ctx, k)
				// 		json.Unmarshal([]byte(val), &m)

				// 		messages = append(messages, m)
				// 	}
				// }
				time.Sleep(time.Second);
				
				if wait++; int(req.GetWaitLimit()) <= wait { break process }

			}

			errChan <- nil 
		}()
		
	return <- errChan
}



func (s *SQSServer) DeleteMessage(ctx context.Context,req *aws.DeleteMessageRequest) (*aws.DeleteResponse,error){
	id := req.GetMessageId()
	queueName := req.GetQueueName()

	msg := model.Message{}

	
	res := s.db.Where("queue_name = ? AND user_message_id = ?",queueName,id).First(&msg)

	if res.Error != nil{
		return &aws.DeleteResponse{Message: "Unable to delete, not found"},nil
	}

	if msg.Status == model.Available || msg.Status == model.Consumed { 
		return &aws.DeleteResponse{Message: "Unable to delete, message is available or already consumed."},nil
	}

	if time.Now().After(msg.DeleteOn) {
		msg.Status = model.Available
		s.db.Save(&msg)
		return &aws.DeleteResponse{Message: "Time to process expired,available on the queue again."},nil
	}


	msg.Status = model.Consumed
	s.db.Save(&msg)
	
	return &aws.DeleteResponse{Message: "Deleted succesfully"},nil
}