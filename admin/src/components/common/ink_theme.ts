import {ref} from "vue";

// 存储当前主题色
export const theme = ref('')
// 设置主题
export const setTheme = (val: string) => {
    if (val === 'dark') {
        document.body.setAttribute('arco-theme', 'dark')
    } else {
        document.body.removeAttribute('arco-theme');
    }
    theme.value = val;
    // 存储主题状态
    localStorage.setItem('theme', val)
}

// 获取当前主题状态
export const loadTheme = () => {
    const val = localStorage.getItem('theme')
    if (val) {
        if (val === 'dark') {
            theme.value = val
            setTheme(val)
        }
    }
}