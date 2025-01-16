/// <reference types="vite/client" />

import "vue-router"

declare  module "vue-router"{
    interface RouterMeta{
        title: string
    }
}

export interface ImportMetaEnv extends  Record<string, string>{
    VITE_SERVER_URL: string
}
