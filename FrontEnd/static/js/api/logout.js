function DeleteHeader() {
    if (document.cookie.includes("token=")) {
        const nav = document.getElementById("nav")
        nav.innerHTML = '<p>Welcome To Forum |</p> <a id="logout">Logout</a>'
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

