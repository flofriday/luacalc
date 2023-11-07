-- Implementations should not directly write to stdout but we should
-- save that somewere for logs etc.
function sum(a, b)
    for _ = 1, 100000 do
        print("I shouldn't be able to write 100_000 lines of text")
    end
    return a + b
end
