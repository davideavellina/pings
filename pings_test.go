/*

Realizzare un programma Go (nome file 'pings.go') che legga un file (il cui nome viene passato come parametro a linea di comando) generato dal comando `fping` (ping multipli su vari host di una rete) e che ne calcoli alcune statistiche:
- quali nodi rispondono
- con che tempistiche (min, max, media totali e per singolo nodo)


Il file contiene righe di due tipi:
1) 192.168.142.136 : [0], 64 bytes, 4.02 ms (4.02 avg, 0% loss)
2) ICMP Host Unreachable from 192.168.142.100 for ICMP Echo sent to 192.168.142.43

Le righe di tipo 1) rappresentano un "successo", il ping ha avuto risposta,
Le righe di tipo 2) rappresentano un "insuccesso" e vanno ignorate.
Le righe di tipo 1) vanno utilizzate per costruire una rappresentazione interna opportuna.

Nel programma deve essere definita:

- una struttura Host contenente i seguenti campi:
	IP string
	pings (slice di float32)
  che verrà utilizzata per le varie elaborazioni.

Vanno definite le seguenti funzioni:

- func (host Host) String() string, che produca una stringa rappresentativa dell'Host host "pingato", nella forma:
	host IP       : lista dei ping times

  Ad esempio:
	192.168.142.242 : 4.62 120 19.2 50.6 2.82 103 195 186 183 48.2

- func addPing(host *Host, pingtime float32)
  che aggiunge un pingtime all'host specificato

- func averageHost(host Host) float32
  che calcola la media dei pingtime dell'Host host

- func minHost(host Host) float32
  che estrae il min pingtime dell'Host host

- func maxHost(host Host) float32
  che estrae il max pingtime dell'Host host

Il programma deve creare un report della forma:
```
...
192.168.142.242 : 4.62 120 19.2 50.6 2.82 103 195 186 183 48.2
192.168.142.243 : nil
...
192.168.142.246 : 10.9 4.46 6.75 9.06 9.18 4.87 9.91 5.84 9.44 6.33
...
```

E in coda deve stampare una statistica nella forma:
```
Numero di host esaminati: ...
Min pingtime: <min>,<host>
Max pingtime: <max>,<host>
Media pingtime: <media>
```

*/

package main

import (
	"fmt"
	"testing"
)

var prog = "./pings"
var diffwidth = 100

func TestMain(t *testing.T) {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(">>> PERCHE' QUESTO TEST FALLIRA' PRATICAMENTE SEMPRE ANCHE SE IL PROGRAMMA E' CORRETTO? <<<")
	fmt.Println()
	fmt.Println()
	fmt.Println()

	ConfrontaConOracolo(
		t,
		prog,
		"",
		"pings.csv") // arg1
}

func TestNoFile(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"sbagliato.input") // arg1
}

func TestNoArg(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"")
}

func TestVuoto(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"vuoto.input") // arg1
}

func TestString(t *testing.T) {
	host := Host{"192.168.1.1", []float32{2.5, 4.7}} // testa anche se la struct è definita
	fmt.Println("String:", host.String())
	if host.String() != "192.168.1.1 : 2.5 4.7" {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	host := Host{"192.168.1.1", []float32{2.5, 4.7}}
	fmt.Println("PRIMA:", host)
	addPing(&host, 3.67)
	fmt.Println("DOPO:", host)
	if len(host.pings) < 3 || host.pings[2] != 3.67 {
		t.Fail()
	}
}

func TestMin(t *testing.T) {
	host := Host{"192.168.1.1", []float32{2.5, 4.7, 2.789, 1}}
	fmt.Println("MIN:", minHost(host))
	if minHost(host) != 1 {
		t.Fail()
	}
}
func TestMax(t *testing.T) {
	host := Host{"192.168.1.1", []float32{2.5, 4.7, 2.789, 1}}
	fmt.Println("MAX:", maxHost(host))
	if maxHost(host) != 4.7 {
		t.Fail()
	}
}

func TestAverage(t *testing.T) {
	host := Host{"192.168.1.1", []float32{2.5, 4.7, 2.789, 1}}
	fmt.Println("AVERAGE:", averageHost(host))
	if averageHost(host)-2.74725 > 10e5 {
		t.Fail()
	}
}
