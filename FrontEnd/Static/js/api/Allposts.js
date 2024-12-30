import { BlackColor , PurpleColor , FormatDate , LikeButton , DislikeButton } from "./Config.js"

let tagfilter
let filters
let offset = 0
const limit = 20 // NOTE: Ensure this value matches the corresponding value in the backend.

const MypostsButton = document.getElementById("myposts")
const LikeFilterButton = document.getElementById("likedposts")
const TopicFilterButton = document.getElementById("topic")
const NextButton = document.getElementById("next")
const PrevionsButton = document.getElementById("previons")

if (NextButton){
    NextButton.addEventListener("click", ()=> {
        offset += limit
        LoadData(filters,offset)
    })
}
if (PrevionsButton){
    PrevionsButton.addEventListener("click", ()=> {
        if (offset > 0 ){
            offset -= limit            
            LoadData(filters,offset)
        }
    })
}

if (MypostsButton) {
    MypostsButton.addEventListener("click", () => {
        LoadData("post")
    })
}
if (LikeFilterButton) {
    LikeFilterButton.addEventListener("click", () => {
        LoadData("like")
    })
}
if (TopicFilterButton) {
    TopicFilterButton.addEventListener("click", () => {
        tagfilter = document.getElementById("tagfilter").value
        LoadData("Tag")
    })
}

async function LoadData(filter="") {
    
    const res = await fetch(`/api/posts?filter=${filter}&offset=${offset}&tagfilter=${tagfilter}`)
    const Data = await res.json()

    if (!Data) {
        handleEmptyDataState()
        return
    }
    if (filter != ""){
        NextButton.disabled = true;
        NextButton.style.display = 'none';
        PrevionsButton.disabled = true
        PrevionsButton.style.display = 'none'
    } else if (NextButton) {
        NextButton.disabled = false;
        NextButton.style.display = 'block'
    }

    const Parent = document.getElementById("forums-container")
    Parent.innerHTML = ""
    Data.forEach(post => {
        let CommentsCounter = 0
        if (post.Comments) {
            CommentsCounter = post.Comments.length
        }


        let Tags = ""
        let LikeIcon = LikeButton(BlackColor , post.Likes.Counter)
        let DislikeIcon = DislikeButton(BlackColor , post.Dislikes.Counter)

        if (post.Likes.IsLiked) {
            LikeIcon = LikeButton(PurpleColor , post.Likes.Counter)
        }
        if (post.Dislikes.IsDislike) {
            DislikeIcon = DislikeButton(PurpleColor , post.Dislikes.Counter)
        }
        if (post.Categories) {
            post.Categories.forEach(category => {
                Tags += `<p>#${category}</p>`
            });
        }
        const Post = document.createElement("div")
        Post.classList.add("forum")
        Post.setAttribute("data-id" , post.Id)

        Post.innerHTML =
            `
            <div class="title">
                <h5 onclick='location.href = "/post?id=${post.Id}"'>${post.Title}</h5>
            </div>
            <div class="content">
                <p>${post.Content.replaceAll("\n", "<br>")}</p>
            </div>
            <div class="topics">
                <div class="tags">
                    ${Tags}
                </div>
                <p>By ${post.UserName} At ${FormatDate(post.CreatedAt)}</p>
            </div>
            <div class="reactions">
                ${LikeIcon}
                ${DislikeIcon}
                <p class="comment" onclick='location.href = "/post?id=${post.Id}"'><svg xmlns="http://www.w3.org/2000/svg" height="20px" viewBox="0 -960 960 960" width="20px" fill="${BlackColor}"><path d="M240-400h320v-80H240v80Zm0-120h480v-80H240v80Zm0-120h480v-80H240v80ZM80-80v-720q0-33 23.5-56.5T160-880h640q33 0 56.5 23.5T880-800v480q0 33-23.5 56.5T800-240H240L80-80Zm126-240h594v-480H160v525l46-45Zm-46 0v-480 480Z"/></svg>${CommentsCounter}</p>
            </div>
            `
        Parent.append(Post)
    });

    const CreatePostEvent = new Event("LoaData")
    document.dispatchEvent(CreatePostEvent)
}

LoadData()
function handleEmptyDataState() {
    if (NextButton) {
        NextButton.disabled = true;
        NextButton.style.display = 'none';
    }  
    const parent = document.getElementById("forums-container");
    parent.innerHTML = "<p>No more posts available</p>";
}
export { LoadData }