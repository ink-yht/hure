export interface RecruiteListType {
    "pagination": {
        "page": number,
        "size": number,
        "total": number,
        "has_next": boolean
    },
    "list": [
        {
            "id": number,
            "company": string,
            "address": string,
            "contact": string,
            "phone": string,
            "created_at": number,
        }
    ]
}