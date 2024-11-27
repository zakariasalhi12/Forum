fetch("/api/posts")
.then((res) => res.json())
.then((res) => {
    LoadData(res)
}
)
.catch(err => console.error(err))

function LoadData(posts) {  
    posts.forEach(posts => {
        
    });

}

