

(()=>{
    const eventData = process.env.EVENT_DATA;
    const handler = process.env.HANDLER;
    
    const path = require("path")
    const ext = path.extname(handler);
    const fn = ext.replace(".","");
    const pt = handler.replace(ext,"");

    const mod = require(pt)

    const code = mod[fn];

    if (typeof code !== "function") {
        console.log(`${fn} is not a function sorry`)
    }

    code(eventData);
    
})()