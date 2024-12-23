import { BlackColor, PurpleColor , FormatDate , LikeButton , DislikeButton } from "./Config.js"

async function PostLoader() {

    const QueryString = window.location.search
    const UrlParam = new URLSearchParams(QueryString)
    const PostID = UrlParam.get("id")

    const res = await fetch(`api/posts?id=${PostID}`)
    let post = await res.json()

    post = post[0]
    const PostContainer = document.getElementById("post-container")
    PostContainer.setAttribute("data-id", post.Id)

    const UserinforReq = await fetch(`api/userinfo?id=${post.User_id}`)
    const UserInfo = await UserinforReq.json()


    let Tags = ""
    let TagsContainer = ""
    let LikeIcon = LikeButton(BlackColor , post.Likes.Counter)
    let DislikeIcon = DislikeButton(BlackColor , post.Dislikes.Counter)

    if (post.Likes.IsLiked) {
        LikeIcon = LikeButton(PurpleColor , post.Likes.Counter)
    }
    if (post.Dislikes.IsDislike) {
        DislikeIcon = LikeButton(PurpleColor , post.Dislikes.Counter)
    }

    if (post.Categories) {
        post.Categories.forEach(category => {
            Tags += `<p>#${category}</p>`
        });
    }

    let CommentsCounter = 0
    if (post.Comments) {
        CommentsCounter = post.Comments.length
    }

    if (Tags != "") {
        TagsContainer = `
        <fieldset class="tags">
            <legend>Tags</legend>
            ${Tags}
        </fieldset>
        `
    }

    PostContainer.innerHTML =
        `
        <div class="post">
            <div class="profile">
                <h3>${post.UserName}</h3>
                    <div>
                        <p>Join Data: ${FormatDate(UserInfo.CreatedAt)}</p>
                    </div>
                    <div>
                        <p>Total Posts: ${UserInfo.TotalPosts}</p>
                    </div>
                    <div>
                        <p>Role: ${UserInfo.Role}</p>
                    </div>
            </div>
            <div class="postContent">
                <div class="postitle">
                    <p>${post.Title}</p>
                </div>
                <div class="postcontent">
                    <p>${post.Content}</p>
                </div>
                ${TagsContainer}
            </div>
        </div>
        <div class="reactions">
            ${LikeIcon}
            ${DislikeIcon}
            <p class="comment"><svg xmlns="http://www.w3.org/2000/svg" height="20px" viewBox="0 -960 960 960" width="20px" fill="${BlackColor}"><path d="M240-400h320v-80H240v80Zm0-120h480v-80H240v80Zm0-120h480v-80H240v80ZM80-80v-720q0-33 23.5-56.5T160-880h640q33 0 56.5 23.5T880-800v480q0 33-23.5 56.5T800-240H240L80-80Zm126-240h594v-480H160v525l46-45Zm-46 0v-480 480Z"/><span>${CommentsCounter}</span></svg></p>
        </div>
    `

    if (CommentsCounter != 0) {
        document.getElementById("comment-container").innerHTML = ''
        
        const Parent = document.getElementById("comment-container")
        post.Comments.forEach(comment => {
            let LikeIcon2 = LikeButton(BlackColor , comment.Likes.Counter , true)
            let DislikeIcon2 = DislikeButton(BlackColor , comment.Dislikes.Counter , true)

            if (comment.Likes.IsLiked) {
                LikeIcon2 = LikeButton(PurpleColor , comment.Likes.Counter , true)
            }
            if (comment.Dislikes.IsDislike) {
                DislikeIcon2 = DislikeButton(PurpleColor , comment.Dislikes.Counter , true)
            }

            const Post = document.createElement("div")
            Post.classList.add("forum")
            Post.setAttribute("data-id", comment.Id)

            Post.innerHTML =
                `
                <div class="title">
                    <h5>${comment.UserName}</h5>
                </div>
                <div class="content">
                    <p>${comment.Content.replaceAll("\n", "<br>")}</p>
                </div>
                <div class="topics">
                    <div class="tags">
                    </div>
                    <p>${FormatDate(comment.CreatedAt)}</p>
                </div>
                <div class="reactions">
                    ${LikeIcon2}
                    ${DislikeIcon2}
                </div>
                `

            Parent.append(Post)
        });
    }

    const CreatePostEvent = new Event("LoaData")
    document.dispatchEvent(CreatePostEvent)
}

PostLoader()

export {PostLoader}