import  {formatDate} from './Allposts'

async function LoadData() {
    const QueryString = window.location.search
    const UrlParam = new URLSearchParams(QueryString)
    const PostID = UrlParam.get("id")

    const res = await fetch("api/posts?id=" + PostID)
    let post = await res.json()
    post = post[0]


    const CreatePostEvent = new Event("LoaData")
    document.dispatchEvent(CreatePostEvent)
}

LoadData()



