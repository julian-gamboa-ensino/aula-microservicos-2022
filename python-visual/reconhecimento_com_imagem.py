import face_recognition
import os
from PIL import Image, ImageDraw

# Diretório local com as imagens de entrada (ajuste conforme necessário)
dir_imagens = "./imagens"

# Diretório para salvar as imagens de saída
dir_saida = "./saida"
if not os.path.exists(dir_saida):
    os.makedirs(dir_saida)  # Cria a pasta "saida" se não existir

# Verifica se a pasta de entrada existe
if not os.path.exists(dir_imagens):
    print(f"A pasta {dir_imagens} não existe. Crie-a e adicione imagens!")
    exit()

# Processa cada arquivo na pasta
for arquivo in os.listdir(dir_imagens):
    caminho_entrada = os.path.join(dir_imagens, arquivo)
    
    try:
        # Carrega a imagem
        imagem = face_recognition.load_image_file(caminho_entrada)
        
        # Detecta rostos
        localizacoes = face_recognition.face_locations(imagem)
        
        if localizacoes:
            print(f"Processado: {arquivo}, {len(localizacoes)} rostos detectados")
            
            # Converte a imagem para o formato PIL para edição
            imagem_pil = Image.fromarray(imagem)
            draw = ImageDraw.Draw(imagem_pil)
            
            # Desenha retângulos ao redor dos rostos
            for (topo, direita, baixo, esquerda) in localizacoes:
                draw.rectangle(((esquerda, topo), (direita, baixo)), outline="red", width=3)
            
            # Salva a imagem com os rostos destacados
            caminho_saida = os.path.join(dir_saida, f"detected_{arquivo}")
            imagem_pil.save(caminho_saida)
            print(f"Imagem salva em: {caminho_saida}")
        else:
            print(f"Nenhum rosto em: {arquivo}")
            
    except Exception as e:
        print(f"Erro ao processar {arquivo}: {e}")