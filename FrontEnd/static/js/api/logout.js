async function Logout() {
    const Res = await fetch("api/isloged");
    const Data = await Res.json();
    if (Res.status == 200) {
        const WelcomSection = document.getElementById("welcome")
        if (WelcomSection) {
            WelcomSection.remove()
        }
        const nav = document.getElementById("nav");
        const cp = document.getElementById("cp")
        if (cp) {
            const CreatePost = `
            <div class="header">
                <p>Create New Post</p>
            </div>
            <div class="fieldsets">
                <div>
                    <b>Post Title :</b>
                    <input id="Posttitle" type="text">
                    </fieldset>
                    <fieldset>
                        <legend>Post Content :</legend>
                        <textarea id="Postcontent" style="resize: none; width: 100%;" rows="5" id="registeremail" ></textarea>
                    </fieldset>
                    <fieldset>
                        <p>Every topic separet by space</p>
                        <legend>Post Topics :</legend>
                        <input id="Postopic" type="text">
                    </fieldset>
                    <button id="createpost">Create Post</button>
                </div>
            </div>`

            cp.innerHTML = CreatePost

            const elementReadyEvent = new Event("elementReady");
            document.dispatchEvent(elementReadyEvent);
        }
        if (nav) {
            nav.innerHTML = `<p>Welcome ${Data.username} |</p> <a id="logout">Logout</a>`;
            sessionStorage.setItem("user_id" , Data.user_id)
        }
        const LogoutButton = document.getElementById("logout");
        if (LogoutButton) {
            LogoutButton.addEventListener("click", async () => {
                await fetch("api/logout", { method: "GET" });
                window.location.href = "/";
            });
        }
        return
    }


    const filters = document.getElementsByClassName("filters")[0]
    if (filters) {
        filters.querySelectorAll("p")[0].remove()
        filters.querySelectorAll("p")[0].remove()
        filters.style.justifyContent = "center"
    }
}

Logout();