import axios from "axios"

export async function apiGet(url, params) {
    await axios.get(url,params).then(res =>{
        // console.log(1)
        // console.log(res.data.data)
        return res.data.data
    }).catch(err =>{
        alert("我们遇到了未知错误，这有可能导致程序无法正常运行");
        return err
    })
}
