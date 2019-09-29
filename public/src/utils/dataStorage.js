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

export function setPermissions(permissionsArr){
    Storage.set(Config.permissionsKey,permissionsArr);
    return permissionsArr
}

export function hasPermissions(permissionsKey){
    let permissionsArr = Storage.get(Config.permissionsKey,[]);
    if(permissionsArr.indexOf('*') !== -1){return true;}
    return permissionsArr.indexOf(permissionsKey) !== -1;
}

export function getToken() {
    return Storage.get(Config.tokenKey)
}

export function setToken(token) {
    return Storage.set(Config.tokenKey, token)
}

export function removeToken() {
    return Storage.remove(Config.tokenKey)
}


export function isLogin() {
    return (getToken() || '').length > 5
}
