document.addEventListener("elementReady", () => {
    const PostButton = document.getElementById("createpost")
    if (PostButton) {
        PostButton.addEventListener("click", async () => {
            const PostTitle = document.getElementById("Posttitle").value
            const PostContent = document.getElementById("Postcontent").value
            const PostTopics = document.getElementById("Postopic").value.split(" ")

            const Res = await fetch("api/post", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    Title: PostTitle,
                    Content: PostContent,
                    Categories: PostTopics,
                }
                )
            })

            if (Res.status != 200) {
                const Data = await Res.json()
                alert(Data.Error)
                return
            }

            alert("Post Created successfuly")
            window.location.href = "/"
        })
    }
})




