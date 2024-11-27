


async function DeleteHeader() {
    const Res = await fetch("api/isloged")
    const Data = await Res.json()
    if (Res.status == 200) {
        const nav = document.getElementById("nav")
        nav.innerHTML = `<p>Welcome ${Data.username} |</p> <a id="logout">Logout</a>`
    }
}

DeleteHeader()

const LogoutButton = document.getElementById("logout")
if (LogoutButton) {
    LogoutButton.addEventListener("click", (e) => {
        fetch("api/logout", {
            method: "DELETE",
        })
        window.location.href = "/"
    })
}
