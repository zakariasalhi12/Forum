
function Register() {
    const RegisterButton = document.getElementById("register")
    if (RegisterButton) {
        RegisterButton.addEventListener('click', async () => {
            const UserName = document.getElementById("registerusername").value
            const email = document.getElementById("registeremail").value
            const password = document.getElementById("registerpassword").value

            const Res = await fetch("api/register", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    username:UserName,
                    email: email,
                    password: password,
                })
            })

            if (Res.status != 200) {
                const Data = await Res.json()
                alert(Data.Error)
                return
            }
            window.location.href = "/"
        })
    }
}

async function CountePosts() {
    const total = document.getElementById("total")
    const Res = await fetch("api/posts")
    const Data = await Res.json()
    total.innerHTML = `Total Posts : <span style="color:var(--Puprle);">${Data.length || 0}</span>`
}

Register()
CountePosts()
