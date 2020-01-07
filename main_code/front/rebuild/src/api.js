import axios from 'axios'

export function apiGet(url,params) {
    axios.get(url,params).then(res =>{
        // console.log(res.data.data)
        return res.data.data
    }).catch(() => {
        alert("我们遇到了网络错误,可能导致程序无法正常运行")
    })
}

export async function apiPost() {

}

export async function apiPut() {

}

export async function apiDelete() {

}
