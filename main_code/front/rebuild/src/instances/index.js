import axios from 'axios'

const ApiErrorMsg = data => {
    const err = data.error;
    switch (err) {
        case 30200: return '您的登录已过期，请重新登录'
        case 40000: return '您填写的格式有错误，请检查后再试'
        case 40010: return '未知错误，请与我们联系'
        case 40020: return '未知错误，请与我们联系'
        case 40030: return '您注册的用户名有重复，请更换用户名再试'
        case 40040: return '未知错误，请与我们联系'
        case 40410: return '您的用户名或密码错误，请检查后登录'
        case 40420: return '您的邀请码错误'
        case 40430: return '您要操作的事务不存在'
        case 50000: return '数据库错误，请与我们联系'
        case 50010: return '未知错误，请与我们联系'
        case 50020: return '未知错误，请与我们联系'
    }
    return `未知错误，请与我们联系（${err}）`
}

const ApiInstance = axios.create();
ApiInstance.interceptors.response.use(
    response => response,
    (error) => {
        alert(ApiErrorMsg(error.response.data))
        console.log(error.response.data)
        return Promise.reject(error.response.data);
    },
);
export { ApiInstance };
