
import { createPinia } from 'pinia';
import persist from 'pinia-plugin-persistedstate';

// 创建大仓库
const pinia = createPinia();
pinia.use(persist);
export default pinia;