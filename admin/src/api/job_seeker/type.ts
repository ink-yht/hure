export interface JobSeekerListType {
    "pagination": {
        "page": number,
        "size": number,
        "total": number,
        "has_next": boolean
    },
    "list": [
        {
            "id": number,
            "user_id": number,
            "name": string,
            "gender": number,
            "phone": string,
            "degree": string,
            "year_of_birth": string,
            "working_hours": string,
            "profile": string,
            "created_at": number,
            "updated_at": number
        }
    ]
}