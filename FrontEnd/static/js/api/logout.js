async function Logout() {
    const Res = await fetch("api/isloged");
    const Data = await Res.json();
    if (Res.status == 200) {
        const nav = document.getElementById("nav");
        if (nav) {
            nav.innerHTML = `<p>Welcome ${Data.username} |</p> <a id="logout">Logout</a>`;
        }
        const LogoutButton = document.getElementById("logout");
        if (LogoutButton) {
            LogoutButton.addEventListener("click", async (e) => {
                await fetch("api/logout", { method: "GET" });
                window.location.href = "/";
            });
        }
    }
}

Logout();