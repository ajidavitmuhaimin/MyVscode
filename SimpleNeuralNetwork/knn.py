import numpy as np

# Data XOR dan labelnya
X = np.array([[0, 0], [0, 1], [1, 0], [1, 1]])
y = np.array([0, 1, 1, 0])

# Fungsi KNN sederhana
def knn_predict(X_train, y_train, X_test, k=1):
    def euclidean_distance(a, b):
        return np.sqrt(np.sum((a - b) ** 2))

    y_pred = []
    for test_point in X_test:
        distances = [euclidean_distance(test_point, x) for x in X_train]
        k_indices = np.argsort(distances)[:k]
        k_nearest_labels = [y_train[i] for i in k_indices]
        y_pred.append(np.argmax(np.bincount(k_nearest_labels)))
    return np.array(y_pred)

# Data yang akan diprediksi
X_test = np.array([[0, 0], [0, 1], [1, 0], [1, 1]])

# Prediksi menggunakan KNN dengan k=1
y_pred = knn_predict(X, y, X_test, k=1)
print("Prediksi:", y_pred)