import {CreateCommentSection} from "./Config.js"

async function CreateComment() {

    const IsLogged = await fetch("/api/isloged")
    if (IsLogged.status != 200) {
        return
    }
    document.getElementById("comment-component").innerHTML = CreateCommentSection
    document.getElementById("createcomment").addEventListener("click", async (e) => {
        const Content = e.target.parentElement.querySelectorAll("textarea")[0].value

        const res = await fetch("api/comment", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                PostId: +document.getElementById("post-container").getAttribute("data-id"),
                Content: Content,
            }
            )
        })

        const Data = await res.json()

        if (res.status != 200) {
            alert(Data.Error)
            return
        }

        alert(Data.Message)
        window.location.href = `/post?id=${document.getElementById("post-container").getAttribute("data-id")}`

    })
}

CreateComment()