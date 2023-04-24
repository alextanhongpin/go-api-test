package contextkey

import "github.com/google/uuid"

const UserID ContextKey[uuid.UUID] = "user_id_ctx"
