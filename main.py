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
'''import tensorflow as tf
import os

cpus=tf.config.experimental.list_physical_devices('CPU')
print(len(cpus))
for cpu in cpus:
    tf.config.experimental.set_memory_growth(cpu,True)

#Usuniemy zdjecie które nie będą z dobrym rozszerzeniem i inne problemy
import cv2 #Komputer vision
import imghdr # Sprawdzanie rozszerzenia

data_dir='testowyfolder'
os.listdir(data_dir)
image_exts =['jpeg','jpg','bmp','png']

for image_class in os.listdir(data_dir): 
    for image in os.listdir(os.path.join(data_dir, image_class)):
        image_path = os.path.join(data_dir, image_class, image)
        try: 
            img = cv2.imread(image_path)
            tip = imghdr.what(image_path)
            if tip not in image_exts: 
                print('Image not in ext list {}'.format(image_path))
                os.remove(image_path)
        except Exception as e: 
            print('Issue with image {}'.format(image_path))'''

import numpy as np
import pandas as pd
import os
import glob
import matplotlib
import seaborn as sns
import matplotlib.pyplot as plt
from wordcloud import WordCloud, STOPWORDS

print(os.listdir("domodelu/ODIR-5K/ODIR-5K"))

data_df = pd.read_excel(open("domodelu/ODIR-5K/ODIR-5K/data.xlsx", 'rb'), sheet_name='Sheet1')



data_df.columns = ["id", 'age', "sex", "left_fundus", "right_fundus", "left_diagnosys", "right_diagnosys", "normal",
                  "diabetes", "glaucoma", "cataract", "amd", "hypertension", "myopia", "other"]

print(data_df)
# !!! To co wyżej to będzie do reprezentacji danych i tylko tyle

# nie wiem jak zrobić model 
# przygotuj tylko std out
# !!! Sprawdzić wypisywanie
import os, glob, cv2
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from sklearn.model_selection import train_test_split
import tensorflow as tf
from tensorflow.keras.layers import *
from tensorflow.keras import backend as K
from tensorflow.keras.models import Sequential
from tensorflow.keras.preprocessing.image import ImageDataGenerator
from tensorflow.keras.utils import get_custom_objects
import efficientnet.tfkeras as efn
from tqdm import tqdm
class Mish(tf.keras.layers.Layer):
    
    def __init__(self, **kwargs):
        super(Mish, self).__init__(**kwargs)
        self.supports_masking = True

    def call(self, inputs):
        return inputs * K.tanh(K.softplus(inputs))

    def get_config(self):
        base_config = super(Mish, self).get_config()
        return dict(list(base_config.items()) + list(config.items()))

    def compute_output_shape(self, input_shape):
        return input_shape
def mish(x):
    return tf.keras.layers.Lambda(lambda x: x*K.tanh(K.softplus(x)))(x)
 
get_custom_objects().update({'mish': Activation(mish)})