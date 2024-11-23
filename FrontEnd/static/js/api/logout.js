document.getElementById("logout").addEventListener("click", (e)=> {
    fetch("api/logout", {
        method:"DELETE",
    })
})