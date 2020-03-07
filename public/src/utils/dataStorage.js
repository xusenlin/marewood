import Storage from 'good-storage'
import Config from '../config/index.js'


export function getUserInfo(key = null){
    let userInfo = Storage.get(Config.userInfoKey);
    if(key)return userInfo.hasOwnProperty(key) ? userInfo[key] : null;
    return userInfo
}

export function setUserInfo(user){
    Storage.set(Config.userInfoKey,user);
    return user
}


export function getToken() {
    return Storage.get(Config.tokenKey)
}

export function setToken(token) {
    return Storage.set(Config.tokenKey, token)
}


export function getSystemInfo(key = null) {
    let systemInfo = Storage.get(Config.systemInfoKey);
    if(key)return systemInfo.hasOwnProperty(key) ? systemInfo[key] : null;
    return systemInfo
}

export function setSystemInfo(info) {
    Storage.set(Config.systemInfoKey,info);
    return info
}

export function removeToken() {
    return Storage.remove(Config.tokenKey)
}

