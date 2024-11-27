
function Login() {
    const loginButton = document.getElementById("login")
    if (loginButton) {
        loginButton.addEventListener('click', async () => {
            const email = document.getElementById("loginemail").value
            const password = document.getElementById("loginpassword").value

            const Res = await fetch("api/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    email: email,
                    password: password,
                })
            })

            if (Res.status != 200) {
                const Data = await Res.json()
                alert(Data.Error)
            }

        })
    }
}

Login()