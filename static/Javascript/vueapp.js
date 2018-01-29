function GetAllTasks(inst) {
    axios.get('/all')
        .then(function (response) {
            console.log(response.data)
            inst.tasks = response.data;
        })
        .catch(function (error) {
            console.log(error)
        })
}

function sub(task) {

    var index = this.tasks.indexOf(task);
    axios.delete('/item/'+task.ID)
    .then(response =>{})
    .catch(e=>{
        this.errors.push(e)
        })
    this.tasks.splice(index, 1);

}


var v1 = new Vue({
    delimiters: ['[[', ']]'],
    el: "#vueApp1",
    data: {
        tasks: [],
        task: '',
    },
    methods: {
        sub: sub
    }

})

GetAllTasks(v1)