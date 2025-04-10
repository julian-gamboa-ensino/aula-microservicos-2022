import face_recognition
import os

# Diretório local com as imagens (ajuste o caminho para sua pasta)
dir_imagens = "./imagens"  # Exemplo: uma pasta chamada "imagens" no mesmo diretório do script

# Verifica se a pasta existe
if not os.path.exists(dir_imagens):
    print(f"A pasta {dir_imagens} não existe. Crie-a e adicione imagens!")
    exit()

# Processa cada arquivo na pasta
for arquivo in os.listdir(dir_imagens):
    caminho = os.path.join(dir_imagens, arquivo)
    
    # Carrega a imagem
    try:
        imagem = face_recognition.load_image_file(caminho)
        
        # Detecta rostos
        localizacoes = face_recognition.face_locations(imagem)
        
        if localizacoes:
            print(f"Processado: {arquivo}, {len(localizacoes)} rostos detectados")
        else:
            print(f"Nenhum rosto em: {arquivo}")
    except Exception as e:
        print(f"Erro ao processar {arquivo}: {e}")