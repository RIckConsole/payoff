# Payoff
## WARNING!!!
Do NOT run this on your own machine, or someone else's machine. Only run this in a controlled virtual environment as this malware does not contain a decrypt function (on purpose). This was made for educational purposes only. 
## What is Payoff?
Payoff is NOT ransomware. Ransomware implies that you can potentially get your data/files back from paying a fee. Payoff does not have any features for decryption or payment of any kind. If you are looking for those features, stay tuned for a future project. 
## What's it do? 
Put simply, it ruins your day. Upon execution, it encrypts the contents of your home directories (documents, desktop, pictures, etc) and  external drives (D:\\, X:\\, etc) using strong AES encryption. It also changes your background to a basic ransom note (but without the ransom! :) )
## Can I send it to a friend as a joke?
NO! As funny as that would be, he would lose access to ALL of his personal files... Permanently!
Though it is possible to run the decrypt function in the FileEncryption library and use the key in the source, STILL DONT RUN IT! The key will not be the same for much longer. 
## Future plans
- randomly generated 32-byte AES key, so you REALLY can't get your data back
- encrypt some Program Files within reason
- Mocking notifications and pop-ups
- Disabling some services