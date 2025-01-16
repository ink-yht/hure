

export function timestampToDateTime(timestamp: number): string {
    // 确保时间戳是以毫秒为单位
    if (isNaN(timestamp) || !isFinite(timestamp)) {
        throw new Error('Invalid timestamp');
    }

    const date = new Date(timestamp);

    // 辅助函数：确保数字为两位数
    function padZero(num: number): string {
        return num.toString().padStart(2, '0');
    }

    const year = date.getFullYear();
    const month = padZero(date.getMonth() + 1);
    const day = padZero(date.getDate());
    const hour = padZero(date.getHours());
    const minute = padZero(date.getMinutes());
    const second = padZero(date.getSeconds());

    return `${year}-${month}-${day} ${hour}:${minute}:${second}`;
}
