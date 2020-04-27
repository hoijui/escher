Flansch, Steck-kupplig, flanschverbindung, (formschluessig,) swage-lock, 

nozzle, faucet, hose, choke


### Gate

* similar like an object, an instance of a class (reflex) in Java(?)
* is a node in the circuit-graph
* has a name (string or int, unique within a circuit)
* has a value (Go type == `interface{}`)

###  Link

* is an (undirected?) edge in the circuit-graph
* a circuit has a set of them
* connects a pair of gates
* has two endpoints called vectors
* also called channel (in the past)?

###  Vector

* has a gate-name
* has a valve-name
* unique withing a circuit (there are no two vectors with the same gate-name & valve-name)

###  Valve

* a specific I/O "whole"/connector
* part of a gate
* has a unique name (withing the gate)
* TODO more info?

###  Reflex

* similar like a class in Java(?)
* may be either of:
** implemented in the underlying technology (Golang)
** composed of gates

### Circuit

* the whole graph?
* the mother reflex/gate?
* TODO


Fabian Bruder, [26.04.19 10:38]
wenn alles strikt deklarativ dorezoge implementiert esch heisst dass jedes objekt emmer typisiert muesst si (er seid dass escher en rein deklarativi sproch esch, aber i ha ned oeberprueft oeb das i de implementierig au so dorezoge esch)

Fabian Bruder, [26.04.19 10:39]
aber jo, s korrekte model waer asynchrons message passing met messages wo typisierti objekt (oder i eusem fall den ev. commands) verschecke






Fabian Bruder, [28.04.19 05:56]

wenni mues en istegsponkt ussueche, de woerdi glaub die membrane (im senn vomene ballon, sphere) nae, wel denn hemmer scho s erschte mapping vo enne flaechi ond osse flaeche definiert...ond i glaub das waer au einigermasse inuitiv zom erklaere.....du hesch 2 gschlossni flaechene, eini of de enne siite, ond eini of de ossesiite, wo sech decke (kongruent send) ond so ennesiite ond ossesiite 1:1 mappe

iez hemmer of dere membrane 2 addressberiich: eine of de ennesite ond eine of de ossere siite.d flaechi saelber wer denn 'root' ond alles was of de flaechi registriert esch woerd denn en tree belde, aehnlech wienes filesystem



s circuit hed met identitaete z tue...das chan es objekt si, es laebewaese oder ergned oebbis dezwoesche (programm, organ, zaelle, verchehrsnetz). wels im noechschte schrett om interaktione goht, definiert die identitaet prmaer en membrane im senn vonere huut wod bezieg ond interaktion definiert ond reguliert. dorom esch of dere ebeni alles komplementaer...wenn niemer de vogel beobachtet, werd de vogel au ned beobachtet. ond wenn de vogel ned beobachtet werd, chas au ke beobaechter gaeh...die existeire emmer metenand ond verschwenede metenand.


 'wahrheits treui' (im senn vo 'high fidelity')


