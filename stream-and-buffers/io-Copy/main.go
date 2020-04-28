//If we want to write only n bytes from an io.Reader to an io.Writer,
//we can use io.CopyN function. This function creates a io.LimitReader internally.