func SaveData1(path string, data []byte) error {
    fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
    if err != nil {
        return err
    }
    defer fp.Close()

    _, err = fp.Write(data)
    if err != nil {
        return err
    }
    return fp.Sync() // fsync
}


# next 10 days busy in building project related to nlp 
# Learnt about mcp servers how it works whether thing
