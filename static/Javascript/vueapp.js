var params = new URLSearchParams();
history.replaceState(null,document.title, location.href);

function GetAllTasks(inst) {
    axios.get('/all')
        .then(function (response) {
            inst.tasks = response.data;
        })
        .catch(function (error) {
            console.log(error)
        })
}

function sub(task) {
    if (confirm('Delete from the list?')) {
        var index = this.tasks.indexOf(task);
    axios.delete('/item/'+task.ID)
    .then(response =>{})
    .catch(e=>{
        this.errors.push(e)
        })
    this.tasks.splice(index, 1);
    } else {

    }
    
}

function populateTf(taskPayload) {

    this.taskBody = taskPayload.Body;
    this.taskCategory = taskPayload.Name;
    this.taskPriority = taskPayload.Priority;
    this.taskID = taskPayload.ID;
    this.editIsActive = "editActive";
    this.currTask = taskPayload;
}

function edit() {
  
    // Have to stringify this to make sure data is sent over in the x-www-form-urlencoded format. Note this method may not work with all browsers, see google. 
    params.delete('taskBody', taskBody);
    params.delete('taskCategory', taskCategory);
    params.delete('taskPriority', taskPriority);
   
    var taskBody = document.getElementsByName("task")[0].value;
    var taskCategory = document.getElementsByName("category")[0].value ;
    var taskPriority = document.getElementsByName("priority")[0].value;
        
    params.append('taskBody', taskBody);
    params.append('taskCategory', taskCategory);
    params.append('taskPriority', taskPriority);

    axios.put('/item/'+this.taskID, params)
    .then(response =>{})
    .catch(e=>{
        this.errors.push(e)
        })
    
    history.replaceState(null,document.title, location.href);
    GetAllTasks(v1);
    this.taskID = 0;
    this.editIsActive = "";
    location.reload();
}

var v1 = new Vue({
    delimiters: ['[[', ']]'],
    el: "#vueApp1",
    data: {
        tasks: [],
        task: '',
        taskBody: '',
        taskCategory: '',
        taskPriority: 1,
        taskID: null,
        currTask: null
    },
    methods: {
        sub: sub,
        edit: edit,
        populateTf: populateTf,
    }

})


GetAllTasks(v1)