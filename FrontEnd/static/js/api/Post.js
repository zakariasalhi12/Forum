const ActiveColor = "#6354bb" // purple

async function PostLoader() {
    const QueryString = window.location.search
    const UrlParam = new URLSearchParams(QueryString)
    const PostID = UrlParam.get("id")

    const res = await fetch(`api/posts?id=${PostID}`)
    let post = await res.json()

    post = post[0]
    const PostContainer = document.getElementById("post-container")
    PostContainer.setAttribute("data-id" , post.Id)

    const UserinforReq = await fetch(`api/userinfo?id=${post.User_id}`)
    const UserInfo = await UserinforReq.json()


    let Tags = ""
    let LikeIcon = `<p class="like"><svg xmlns="http://www.w3.org/2000/svg" height="20px" viewBox="0 -960 960 960" width="20px" fill="#222"><path d="M720-120H280v-520l280-280 50 50q7 7 11.5 19t4.5 23v14l-44 174h258q32 0 56 24t24 56v80q0 7-2 15t-4 15L794-168q-9 20-30 34t-44 14Zm-360-80h360l120-280v-80H480l54-220-174 174v406Zm0-406v406-406Zm-80-34v80H160v360h120v80H80v-520h200Z"/></svg><span>${post.Likes.Counter}</span></p>`
    let DislikeIcon = `<p class="dislike"><svg xmlns="http://www.w3.org/2000/svg" height="20px" viewBox="0 -960 960 960" width="20px" fill="#222"><path d="M240-840h440v520L400-40l-50-50q-7-7-11.5-19t-4.5-23v-14l44-174H120q-32 0-56-24t-24-56v-80q0-7 2-15t4-15l120-282q9-20 30-34t44-14Zm360 80H240L120-480v80h360l-54 220 174-174v-406Zm0 406v-406 406Zm80 34v-80h120v-360H680v-80h200v520H680Z"/></svg><span>${post.Dislikes.Counter}</span></p>`

    if (post.Likes.IsLiked) {
        LikeIcon = `<p class="like" style="color:${ActiveColor};"><svg xmlns="http://www.w3.org/2000/svg" height="20px" viewBox="0 -960 960 960" width="20px" fill="#6354bb"><path d="M720-120H280v-520l280-280 50 50q7 7 11.5 19t4.5 23v14l-44 174h258q32 0 56 24t24 56v80q0 7-2 15t-4 15L794-168q-9 20-30 34t-44 14Zm-360-80h360l120-280v-80H480l54-220-174 174v406Zm0-406v406-406Zm-80-34v80H160v360h120v80H80v-520h200Z"/></svg><span>${post.Likes.Counter}</span></p>`
    }
    if (post.Dislikes.IsDislike) {
        DislikeIcon = `<p class="dislike" style="color:${ActiveColor};"><svg xmlns="http://www.w3.org/2000/svg" height="20px" viewBox="0 -960 960 960" width="20px" fill="#6354bb"><path d="M240-840h440v520L400-40l-50-50q-7-7-11.5-19t-4.5-23v-14l44-174H120q-32 0-56-24t-24-56v-80q0-7 2-15t4-15l120-282q9-20 30-34t44-14Zm360 80H240L120-480v80h360l-54 220 174-174v-406Zm0 406v-406 406Zm80 34v-80h120v-360H680v-80h200v520H680Z"/></svg><span>${post.Dislikes.Counter}</span></p>`
    }
    post.Categories.forEach(category => {
        Tags += `<p>#${category}</p>`
    });

    PostContainer.innerHTML = 
    `
                    <div class="post">
                    <div class="profile">
                        <h3>${post.UserName}</h3>
                        <div>
                            <p>Join Data: ${UserInfo.CreateDate}</p>
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

                        <fieldset class="tags">
                            <legend>Tags</legend>
                            ${Tags}
                        </fieldset>
                    </div>

                </div>
                <div class="reactions">
                    ${LikeIcon}
                    ${DislikeIcon}
                    <p class="comment"><svg xmlns="http://www.w3.org/2000/svg" height="20px" viewBox="0 -960 960 960" width="20px" fill="#222"><path d="M240-400h320v-80H240v80Zm0-120h480v-80H240v80Zm0-120h480v-80H240v80ZM80-80v-720q0-33 23.5-56.5T160-880h640q33 0 56.5 23.5T880-800v480q0 33-23.5 56.5T800-240H240L80-80Zm126-240h594v-480H160v525l46-45Zm-46 0v-480 480Z"/></svg></p>
                </div>
    `



    const CreatePostEvent = new Event("LoaData")
    document.dispatchEvent(CreatePostEvent)
}

PostLoader()



