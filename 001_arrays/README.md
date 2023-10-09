# Datenstrukturen 1

## If you get lost

- https://simpleclub.com/lessons/fachinformatikerin-array
- https://de.wikipedia.org/wiki/Array_(Datentyp)
- https://de.wikipedia.org/wiki/Schach
- https://zetcode.com/golang/file

## Aufgabe

### Einleitung

Du bist ein Schulleiter im Jahre 1999 und das alljährliche Schachturnier steht
vor der Tür. Es haben sich so viele Teilnehmer angemeldet wie noch nie und 10
Leute die als Aufseher eingeplant waren haben bereits abgesagt.
Kein Problem du warst schlau und hast vorgesorgt. Apple Newtons sind der neueste
Schrei und mit diesen futuristischen PDAs können die Teilnehmer ihr Spiel selbst
festhalten, damit die Auswertung in Ruhe nacheinander gemacht werden kann.

Leider gibt es da dann doch ein paar Probleme ...

### Aufgabe 1

Das Programm das zum protokollieren des Spiels installiert sein sollte fehlt
komplett auf allen Newtons. Da muss der Chef wohl wieder selbst ran ...


Folgende Boilerplate hast du noch auf deinem C64 gehabt in einer Sprache die dir
dein guter Freund Ken Thompson zum probieren auf Diskette gegeben hat. Du hast
sogar schon angefangen, was für ein Einstieg. Das kann nur ein Erfolg werden!
```go
package main

import (
 "flag"
 "fmt"
)

func printf(msg string, values...interface{}) {
  str := fmt.Sprintf(msg, values...)
  fmt.Println(str)
}

func main() {
  piece := flag.String("piece", "", "The chess piece to move")
  pos := flag.String("pos", "", "The position where to put the chess piece FROM:TO")

  flag.Parse()

  printf("%v, %v", *piece, *pos)
}
```

 - a) Schreibe eine simple Konsolen-Anwendung (eine CLI) der man eine Figur und
      eine Feldbezeichnung übergeben kann und die das aktualisierte Schachbrett
      in einer Textdatei speichert.
    - Das Schachbrett soll mit der Standardaufstellung initialisiert werden.
      Folgendes Muster ist gegeben (die Backslashes gehören nicht dazu):

      sT, sS, sL, sK, sD, sL, sS, sT \
      sB, sB, sB, sB, sB, sB, sB, sB \
       e,  e,  e,  e,  e,  e,  e,  e \
       e,  e,  e,  e,  e,  e,  e,  e \
       e,  e,  e,  e,  e,  e,  e,  e \
       e,  e,  e,  e,  e,  e,  e,  e \
      wB, wB, wB, wB, wB, wB, wB, wB \
      wT, wS, wL, wD, wK, wL, wS, wT

    - Wird eine Position geändert soll vorher überprüft werden ob dieser Zug
      auch möglich ist.


