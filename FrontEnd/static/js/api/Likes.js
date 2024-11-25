document.getElementById("createpost").addEventListener("click", (e) => {
    fetch("api/like", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            PostOrCommentId: 1,
            IsComment: true,
            IsLike: true,
        }
    )
    })
})

