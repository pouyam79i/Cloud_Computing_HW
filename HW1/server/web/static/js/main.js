// Import files in html
function importText(eID, type) {
    var data = null
    let input = document.createElement('input');
    input.type = 'file';
    input.multiple = false;
    input.accept = type;
    input.onchange = _ => {
      // you can use this method to get file and perform respective operations
            var fr = new FileReader();
            fr.onload=function(){
                displayItem(eID, fr.result);
            }
            fr.readAsText(input.files[0]);

        };
    input.click();
}

function displayItem(eID, text){
    console.log("Loading data on ID:",eID,"\n", text);
    document.getElementById(eID).innerText = `${text}`;
}

// login form action here:
function singInUP() {

    console.log("Submit clicked")

    // reading email password
    let username = document.getElementById("email").value.trim().toLowerCase();
    let password = document.getElementById("psw").value.trim();

    if (username == "" || password == "") {
       console.log("Plz insert email:password properly")
    } else {
        if (!username.endsWith("@gmail.com")){
            console.log("Use a gmail to sign in")
            alert("Use a gmail to sign in!")
            return
        }
        if (username.length > 512){
            alert("Gmail with more than 512 char is not valid!")
            return
        }
        if (password.length > 512){
            alert("Password with more than 512 char is not valid!")
            return
        }
        console.log("Email:", username, "Password:", password)
        // TODO: post data to server
        try{
            postSingInData(username, password);
        }catch(e){
            console.log("Failed to Post Data!")
        }
    } 
}

// Data sender
async function postSingInData(username, password){
    let response = await fetch("http://localhost:8085/signin", {
    method: 'POST',
    headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
    'Access-Control-Allow-Origin' : '*', 
    'Access-Control-Allow-Credentials' : true 
    },
    body: `{
    "email": "${username}",
    "password": "${password}"
    }`,
    }).then(response=>response.json())
    .then(data=>{ console.log(data); });
}