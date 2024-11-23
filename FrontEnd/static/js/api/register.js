document.getElementById("register").addEventListener("click", (e) => {
    const UserName = e.target.parentElement.querySelectorAll("input")[0].value
    const Email = e.target.parentElement.querySelectorAll("input")[1].value
    const Password = e.target.parentElement.querySelectorAll("input")[2].value

    fetch("api/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            username: UserName,
            email: Email,
            password: Password,
        }
    )
    })
})

