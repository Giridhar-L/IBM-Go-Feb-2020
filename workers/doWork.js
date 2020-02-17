function doWork() {
    for (let i = 0; i < 30000; i++) {
        if (i % 300 === 0){
            var percentCompleted = (i / 30000) * 100;
            self.postMessage({ type : 'PROGRESS', percentCompleted : percentCompleted})
        }
        for (let j = 0; j < 10000; j++) {
            for (let k = 0; k < 100; k++) {


            }

        }

    }
}

self.addEventListener('message', (evt)=>{
    if (evt.data.type === 'START'){
        doWork();
        self.postMessage({ type : 'DONE'});
    } else {
        console.log('unknown message ', evt.data);
    }
})