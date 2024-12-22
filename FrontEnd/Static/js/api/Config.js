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

function FormatDate(date) {
    const day = new Date(date)
    const month = day.getMonth() + 1
    const currentDay = day.getDate()
    const year = day.getFullYear()
    return `${month}/${currentDay}/${year}`;
}

export{PurpleColor , BlackColor , CreateCommentSection , CreatePostSection, FormatDate}