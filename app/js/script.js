window.onload = function(){
    var btn = document.getElementById("btn");
    btn.addEventListener("click",function(){
        send("POST","/post",JSON.stringify({ "ajax_post_data" : 'hello'}));
    });

    var taskInput = document.getElementById("task");
    taskInput.addEventListener("keydown",function(event){
        if (taskInput.value.trim() != "" && event.keyCode == 13){
            console.log(taskInput.value.trim())
        }
    })
};

function send(method, url, data) {
    if(!data){
        data = null
    }
    var request = new XMLHttpRequest();
    request.open(method, url, true);
    request.setRequestHeader("Content-type", "application/json");
    request.send(data)
}