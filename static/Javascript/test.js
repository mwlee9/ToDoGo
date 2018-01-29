function GetAllTasks(inst) {
    axios.get('/all')
    .then(function(response) {
        console.log(response.data)
        inst.parlist = response.data;
    })
    .catch(function(error) {
        console.log(error)
    })
}

var vuecomp = Vue.component("tasklist", {
    delimiters: ['[[', ']]'],
    props: ["childlist"],
    template: `
    <div> 
        <li v-for="child in childlist">[[child]]</li>
    </div>
    `,
   


    data: function() {
        return {childlist1: [123], task: 'Hello'}
    }

});



var v1 = new Vue({
    delimiters: ['[[', ']]'],
    el: "#vueApp1",
    
    data: {
        parlist: [],
        newTask: {},
        mes: "Hello"
    }
    
})

GetAllTasks(v1)