async function LoadData() {
    const res = await fetch("api/posts")
    const Data = await res.json()
    console.log(Data)
}

LoadData()



