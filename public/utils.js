function secsToMinSecs(secs) {
    let minutes = Math.floor(secs / 60);
    let rv = minutes == 0 ? "" : (minutes + "分");
    let seconds = secs - minutes * 60;
    rv += ((seconds < 10 ? "0" : "") + seconds + "秒");

    return rv;
}

Date.prototype.format = function(fmt) {
    let ret;
    const opt = {
        "Y+": this.getFullYear().toString(), // 年
        "m+": (this.getMonth() + 1).toString(), // 月
        "d+": this.getDate().toString(), // 日
        "H+": this.getHours().toString(), // 时
        "M+": this.getMinutes().toString(), // 分
        "S+": this.getSeconds().toString() // 秒
        // 有其他格式化字符需求可以继续添加，必须转化成字符串
    };

    for (let k in opt) {
        ret = new RegExp("(" + k + ")").exec(fmt);
        if (ret) {
            fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
        };
    };
    return fmt;
}

function sureCall(f) {
    if (typeof f === "function") { f(); }
}