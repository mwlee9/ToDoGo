function getHighPayload() {
    axios.get('/highPriority', highPriorityResponse)
    .then(response =>{})
    .catch(e=>{
        this.errors.push(e)
        })
    console.log(highPriorityResponse);
    console.log("TEST");

}


function add(task) {
    return this.list.push(task);

}

function sub(task) {
    var index = this.list.indexOf(task);
    this.list.splice(index, 1);
    // console.log(typeof(index))
    axios.delete('/'+index.toString())
    .then(response =>{})
    .catch(e=>{
        this.errors.push(e)
        })
  }


function savePayload(valRec) {
        axios.post('/save', valRec)
        .then(response =>{})
        .catch(e=>{
            this.errors.push(e)
        })
    }
    
function sendPayload(eventML) {
    this.$emit('tasklistPayload', {tasklist: this.list, compId: this._uid})
}

var vuecomp = Vue.component("tasklist", {
    template: `
    <div> 
    <input v-model="task" type="text">
    <button @click="add(task)">+</button>
    <li v-for="task in list" @click="sub(task)">{{task}}</li>
    <button @click="sendPayload">SEND</button>
    </div>
    `,
    

    data: 
    function() {
        return {
            list: [], task: '',  
        }
    },

   
    
    methods: {
        add: add,
        sub: sub,
        sendPayload: sendPayload

    } 

});

Vue.component('root',{

    template: `
    <div>
        <tasklist @tasklistPayload="savePayload($event)"></tasklist>
    </div>
    `,

    methods: {
        savePayload: savePayload
    }

})


var v1 = new Vue({

    el: "#vueApp1",
       
});

new Vue({
    el: "#vueApp2",
});

new Vue({
    el: "#vueApp3",
});
