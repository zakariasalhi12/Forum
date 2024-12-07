const ActiveColor = "#6354bb" // purple

function Like_DisLike() {

    document.addEventListener("LoaData", () => {
        document.querySelectorAll(".like").forEach(button => {
            button.addEventListener("click", () => {
                Likeer(true, false, button, button.parentElement.querySelector(".dislike"))
            })
        })
        document.querySelectorAll(".dislike").forEach(button => {
            button.addEventListener("click", () => {
                Likeer(false, false, button.parentElement.querySelector(".like"), button)
            })
        })
        document.querySelectorAll(".clike").forEach(button => {
            button.addEventListener("click", () => {
                Likeer(true, true, button, button.parentElement.querySelector(".cdislike"))
            })
        })
        document.querySelectorAll(".cdislike").forEach(button => {
            button.addEventListener("click", () => {
                Likeer(false, true, button.parentElement.querySelector(".like"), button)
            })
        })

    })

}

async function Likeer(Islike, isComment, LikeButton, DislikeButton) {
    let Id;

    if (isComment) {
        Id = +document.getElementById("post-container").getAttribute("data-id")
    }else {
        Id = +LikeButton.parentElement.parentElement.getAttribute("data-id")
    }

    const Res = await fetch("api/like", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            PostOrCommentId: +Id,
            IsComment: isComment,
            IsLike: Islike,
        }
        )
    })

    if (Res.status != 200) {
        window.location.href = "/register"
        return
    }

    const Data = await Res.json()

    if (Data.AlreadyLiked) {
        Like_DisLike_Dom_Handler("#222" , LikeButton , Data.PostsLikes)
        Like_DisLike_Dom_Handler("#222" , DislikeButton , Data.PostsDislikes)
        return
    }
    if (Islike) {
        Like_DisLike_Dom_Handler(ActiveColor , LikeButton , Data.PostsLikes)
        Like_DisLike_Dom_Handler("#222" , DislikeButton , Data.PostsDislikes)
        return
    }
    Like_DisLike_Dom_Handler("#222" , LikeButton , Data.PostsLikes)
    Like_DisLike_Dom_Handler(ActiveColor , DislikeButton , Data.PostsDislikes)
}

function Like_DisLike_Dom_Handler(color , button , data) {
    button.style.color = color
    button.querySelector("svg").setAttribute("fill" , color)
    button.querySelector("span").innerHTML = data
}

Like_DisLike()

