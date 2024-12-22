const PurpleColor = "#6354bb"
const BlackColor = "#222"

const CreateCommentSection =
`
<div class="header">
    <p>Create New Comment</p>
</div>
<div class="fieldsets">
<div>
    <fieldset>
        <legend>Comment Content :</legend>
        <textarea id="commentcontent" style="resize: none; width: 100%;" rows="5" id="registeremail" ></textarea>
    </fieldset>
    <button id="createcomment">Create Comment</button>
</div>
</div>
`
const CreatePostSection = 
`
<div class="header">
    <p>Create New Post</p>
</div>
<div class="fieldsets">
    <div>
        <b>Post Title :</b>
        <input id="Posttitle" type="text">
        <fieldset>
            <legend>Post Content :</legend>
            <textarea id="Postcontent" style="resize: none; width: 100%;" rows="5" id="registeremail" ></textarea>
        </fieldset>
        <fieldset>
            <p>Every topic separet by space . The maximum number of topics is 6</p>
            <legend>Post Topics :</legend>
            <input id="Postopic" type="text">
        </fieldset>
        <button id="createpost">Create Post</button>
    </div>
</div>
`

function LikeButton(color , data) {
    return `<p class="like" style="color:${color};"><svg xmlns="http://www.w3.org/2000/svg" height="20px" viewBox="0 -960 960 960" width="20px" fill="${color}"><path d="M720-120H280v-520l280-280 50 50q7 7 11.5 19t4.5 23v14l-44 174h258q32 0 56 24t24 56v80q0 7-2 15t-4 15L794-168q-9 20-30 34t-44 14Zm-360-80h360l120-280v-80H480l54-220-174 174v406Zm0-406v406-406Zm-80-34v80H160v360h120v80H80v-520h200Z"/></svg><span>${data | 0}</span></p>`
}

function DislikeButton(color , data) {
    return `<p class="dislike" style="color:${color};"><svg xmlns="http://www.w3.org/2000/svg" height="20px" viewBox="0 -960 960 960" width="20px" fill="${color}"><path d="M240-840h440v520L400-40l-50-50q-7-7-11.5-19t-4.5-23v-14l44-174H120q-32 0-56-24t-24-56v-80q0-7 2-15t4-15l120-282q9-20 30-34t44-14Zm360 80H240L120-480v80h360l-54 220 174-174v-406Zm0 406v-406 406Zm80 34v-80h120v-360H680v-80h200v520H680Z"/></svg><span>${data | 0}</span></p>`
}

function FormatDate(date) {
    const day = new Date(date)
    const month = day.getMonth() + 1
    const currentDay = day.getDate()
    const year = day.getFullYear()
    return `${month}/${currentDay}/${year}`;
}


export{PurpleColor , BlackColor , CreateCommentSection , CreatePostSection, FormatDate , LikeButton , DislikeButton}