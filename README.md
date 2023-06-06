### ROT Simple Encryption Algorithm for Chinese and English, Numbers

### Introduction
Inspired by [ROT13](https://en.wikipedia.org/wiki/ROT13), I implemented a simple ROT encryption algorithm that can encrypt Chinese, English, and numbers, but does not support special characters.
The current Chinese character set used contains 4846 commonly used Chinese characters. The character set can be customized by modifying the `words` variable in `rot_han.go`. Other character sets include 26 common English letters and 10 digits.

### How to use       
```go
// Encrypt
ROT("Hello World!") // "Uryyb Jbeyq!"    
ROT("你好，世界！") // "匾竟，粹鸯！"      
ROT("2023-06-06 苹果公司(Apple, AAPL)发布了Vision Pro头戴式设备") // "7578-51-51 祠址夆季(Nccyr, NNCY)呆报盥Ivfvba Ceb苯漪悌揆庙"    

// Decrypt
ROT("Uryyb Jbeyq!") // "Hello World!"      
ROT("匾竟，粹鸯！") // "你好，世界！" 
ROT("7578-51-51 祠址夆季(Nccyr, NNCY)呆报盥Ivfvba Ceb苯漪悌揆庙") // ""2023-06-06 苹果公司(Apple, AAPL)发布了Vision Pro头戴式设备"    
```    
### To do list
- [ ] To support custom rotation offsets. The rotation offset for each character set can be specified each time encryption is performed, increasing the difficulty of cracking the encryption.

### Attention   
The ROT encryption algorithm is a relatively simple algorithm, not suitable for encrypting important information, for learning and reference only.    

### License   
MIT
