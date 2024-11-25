document.getElementById("createpost").addEventListener("click", (e) => {
    fetch("api/like", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            PostId: 1,
            IsComment: false,
            IsLike: false,
        }
    )
    })
})

