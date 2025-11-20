import heapq

def dijkstra(grafo, sorgente):
    # Distanze iniziali: inf per tutti, 0 per la sorgente
    dist = {nodo: float('inf') for nodo in grafo}
    dist[sorgente] = 0
    
    # Coda di priorità (distanza, nodo)
    coda = [(0, sorgente)]
    
    # Per ricostruire il percorso (opzionale)
    predecessore = {nodo: None for nodo in grafo}
    
    while coda:
        distanza_attuale, nodo_attuale = heapq.heappop(coda)
        
        # Se prendiamo un nodo con distanza maggiore di quella registrata, lo ignoriamo
        if distanza_attuale > dist[nodo_attuale]:
            continue
        
        # Esplora i vicini
        for vicino, peso in grafo[nodo_attuale]:
            nuova_dist = distanza_attuale + peso
            
            # Se abbiamo trovato un percorso migliore
            if nuova_dist < dist[vicino]:
                dist[vicino] = nuova_dist
                predecessore[vicino] = nodo_attuale
                heapq.heappush(coda, (nuova_dist, vicino))
    
    return dist, predecessore


# Esempio di grafo (puoi provarlo)
if __name__ == "__main__":
    grafo = {
        'S': [('A', 4), ('B', 2), ('C', 5)],
        'A': [('D', 10), ('S', 4)],
        'B': [('C', 3), ('E', 3), ('S', 2)],
        'C': [('B', 3), ('D', 1), ('S', 5)],
        'D': [('A', 10), ('C', 1), ('E', 2)],
        'E': [('B', 3), ('D', 2)],
    }

    distanze, pred = dijkstra(grafo, 'S')
    print("Distanze minime:", distanze)
    print("Predecessori:", pred)
