<!-- PROJECT LOGO -->
<p align="center">
  <img src="./img/hangman.png" />
</p>
<br />
<div align="center">
    

  <h1 align="center">HangMan Web project build with HTML CSS and Golang</h1>

  <p align="center">
    <br />
  </p>
</div>



<!-- ABOUT THE PROJECT -->
### Built With

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/768px-Go_Logo_Blue.svg.png?20191207190041" alt="golang" width="200">
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/6/61/HTML5_logo_and_wordmark.svg/768px-HTML5_logo_and_wordmark.svg.png?20170517184425" alt="golang" width="200">
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/d/d5/CSS3_logo_and_wordmark.svg/544px-CSS3_logo_and_wordmark.svg.png" alt="golang" width="160">

<br>
<br>

We Have used Golang for Back-end (data traitment) , and HTML & CSS for front-end (design, etc)

</br>

<!-- GETTING STARTED -->
## Getting Started

Here we have created a golang version of the famous hangman game using web interface. You have 10 attempts to find a word, some letters will be revealed at the start, the objective is to find them all in order to win the game.

### Features available

<ul>
  <li>Propose one letter or one word</li>
  <li>Screen Win-Loose</li>
  <li>Hangman draw</li>
  <li>Reset Button</li>
  <li>DataSave to a text file</li>
</ul>

<br>
You can propose one letter or one word if the letter are not present in the word you loose one Attempts and if the word propose is wrong you loose two attempts.
<br>
<br>
If you win you will be redirected to a page where you can start playing again. If you loose you will be redirected to a page where you can start playing again.
<br>
<br>

If the proposal is false you will lose an attempt and the drawing of the hanged man will be drawn as you go.
<br>
<br>
If you find that the proposed word is too complicated to find you can start the match again without losing it
<br>
<br>
When you win the game all the informations about the games are saved on score.txt
<br>
<br>

## Prerequisites
You need to have the go language installed beforehand

[Golang](https://go.dev/dl/)

### Installation

_To use the project you should clone the repo._

Clone the repo
   ```sh
   git clone https://ytrack.learn.ynov.com/git/lmatheo/hangman-web.git
   ```


<!-- RUN THE PROJECT -->
### Run the project
```
cd hangman-web
```
</br>

```go
go run ./server/main.go
```

Then open your navigator and go to localhost:8080, enter username and enjoy.


### Commands in game 


##### Submit one letter or one word
##### Reset game if you want to change word
