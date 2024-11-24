document.getElementById("createpost").addEventListener("click", (e) => {
    const Title = e.target.parentElement.querySelectorAll("input")[0].value
    const Content = e.target.parentElement.querySelectorAll("input")[1].value
    const categories = e.target.parentElement.querySelectorAll("input")[2].value

    fetch("api/newpost", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            Title: Title,
            Content: Content,
            Categories: categories.split(" "),
        }
    )
    })
})

