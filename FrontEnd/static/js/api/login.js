document.getElementById("login").addEventListener("click", (e) => {
    const Email = e.target.parentElement.querySelectorAll("input")[0].value
    const Password = e.target.parentElement.querySelectorAll("input")[1].value

    fetch("api/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            Email: Email,
            Password: Password,
        })
    })
})