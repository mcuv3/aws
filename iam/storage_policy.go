package iam

import aws "github.com/MauricioAntonioMartinez/aws/proto"

func (storage *PostgreStorage) CreatePolicy(in *aws.Policy) (*aws.Policy, error) {
	stmt, err := storage.db.Prepare(`
		INSERT INTO policys(
			name,
			manifest,
			accountId
		)
		VALUES ($1,$2,$3)
		RETURNING id;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		 in.Name,
		 in.Manifest,
		 in.AccountId,
	).Scan(
		&(in.Id),
	)

	return in, err
}


func (storage *PostgreStorage) DeletePolicy(id uint32) error {
	stmt, err := storage.db.Prepare("DELETE FROM policys WHERE id=$1;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}


func (storage *PostgreStorage) GetPolicy(id uint32) (*aws.Policy, error) {
	stmt, err := storage.db.Prepare(`
		SELECT
			id,
			name,
			manifest,
			accountId
		FROM policys
		WHERE id=$1;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	ret := &aws.Policy{}
	err = stmt.QueryRow(id).Scan(
		&ret.Id,
		&ret.Name,
		&ret.Manifest,
		&ret.AccountId,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}


func (storage *PostgreStorage) UpdatePolicy(updated *aws.Policy) (*aws.Policy, error) {
	tx, err := storage.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`
		UPDATE policys
		SET
			name=$1,
			manifest=$2,
			accountId=$3
		WHERE
			id=$5
		RETURNING
			id,
			name,
			manifest,
			accountId
		;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	ret := &aws.Policy{}
	err = stmt.QueryRow(
		updated.Name,
		updated.Manifest,
		updated.AccountId,
	).Scan(
		&ret.Id,
		&ret.Name,
		&ret.Manifest,
		&ret.AccountId,
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	return ret, err
}
