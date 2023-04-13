# Można stworzyć model A.I. do diagnozowania chorób oka na podstawie zdjęć oka, wykorzystując sieci neuronowe konwolucyjne (Convolutional Neural Networks - CNN). Model taki byłby w stanie nauczyć się rozpoznawać różne choroby oka, takie jak zaćma, jaskra czy zwyrodnienie plamki żółtej.

# Do stworzenia takiego modelu potrzebna jest duża ilość danych, czyli zdjęć oka w różnych pozycjach i różnych przypadkach chorobowych. Można skorzystać z publicznych zbiorów danych, takich jak OCT-ImageNet, EyePACS lub Fovea-Net, a także zebrać własny zbiór danych.

# Model A.I. powinien działać w następujący sposób:

# Otrzymuje zdjęcie oka jako wejście.
# Przetwarza zdjęcie za pomocą sieci neuronowej konwolucyjnej w celu rozpoznania chorób oka.
# Zwraca wynik procentowy dla każdej z chorób oka, które zostały rozpoznane, wraz z etykietami diagnoz.
# Wynik można wyświetlić w standardowym wyjściu (stdout) lub w interfejsie użytkownika. W przypadku, gdy wynik procentowy dla danej choroby przekroczy pewien próg, można wskazać, że dana osoba jest chora. Jednak zawsze należy pamiętać, że diagnozowanie chorób powinno być przeprowadzane przez wykwalifikowanego specjalistę, a wynik modelu A.I. należy traktować tylko jako wskazówkę.


# ! https://www.youtube.com/watch?v=jztwpsIzEGc&t=186s 



'''
będzie potrzebne jak deploy na cloud
pip install tenserflow tensorflow-gpu opencv-python matplotlib

'''